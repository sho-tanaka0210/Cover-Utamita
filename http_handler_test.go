package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestHealthCheckHasNoSideEffects(t *testing.T) {
	var runs atomic.Int32
	handler := newHTTPHandler(newRunGuard(t.TempDir()), func(context.Context) error {
		runs.Add(1)
		return nil
	}, nil)

	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if got := runs.Load(); got != 0 {
		t.Fatalf("runner called %d times, want 0", got)
	}
}

func TestRunRejectsUnauthenticatedRequestBeforeExecution(t *testing.T) {
	var runs atomic.Int32
	handler := newHTTPHandler(newRunGuard(t.TempDir()), func(context.Context) error {
		runs.Add(1)
		return nil
	}, nil)

	request := httptest.NewRequest(http.MethodPost, "/run", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusUnauthorized)
	}
	if got := runs.Load(); got != 0 {
		t.Fatalf("runner called %d times, want 0", got)
	}
}

func TestRunOnlyAcceptsPost(t *testing.T) {
	handler := newHTTPHandler(newRunGuard(t.TempDir()), func(context.Context) error {
		t.Fatal("runner must not be called")
		return nil
	}, nil)

	request := httptest.NewRequest(http.MethodGet, "/run", nil)
	request.Header.Set("Authorization", "Bearer token")
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusMethodNotAllowed)
	}
}

func TestRunExecutesTargetDateOnlyOnce(t *testing.T) {
	var runs atomic.Int32
	now := func() time.Time {
		return time.Date(2026, time.August, 9, 12, 0, 0, 0, jst())
	}
	handler := newHTTPHandler(newRunGuard(t.TempDir()), func(context.Context) error {
		runs.Add(1)
		return nil
	}, now)

	for range 2 {
		request := httptest.NewRequest(http.MethodPost, "/run", nil)
		request.Header.Set("Authorization", "Bearer token")
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
		}
	}

	if got := runs.Load(); got != 1 {
		t.Fatalf("runner called %d times, want 1", got)
	}
}

func TestRunGuardSerializesConcurrentRequests(t *testing.T) {
	guard := newRunGuard(t.TempDir())
	started := make(chan struct{})
	release := make(chan struct{})
	var runs atomic.Int32
	run := func() error {
		if runs.Add(1) == 1 {
			close(started)
		}
		<-release
		return nil
	}

	results := make(chan bool, 2)
	errors := make(chan error, 2)
	go func() {
		executed, err := guard.Run("2026-08-08", run)
		results <- executed
		errors <- err
	}()
	<-started
	go func() {
		executed, err := guard.Run("2026-08-08", run)
		results <- executed
		errors <- err
	}()
	close(release)

	first, second := <-results, <-results
	if err := <-errors; err != nil {
		t.Fatal(err)
	}
	if err := <-errors; err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatalf("executed results = %v, %v; want one true and one false", first, second)
	}
	if got := runs.Load(); got != 1 {
		t.Fatalf("runner called %d times, want 1", got)
	}
}
