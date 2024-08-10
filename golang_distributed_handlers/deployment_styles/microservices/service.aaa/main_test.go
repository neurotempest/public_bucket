package main

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestRun(t *testing.T) {

	testCases := []struct{
		Method string
		Endpoint string
		ReqBody string
		ExpectedRespStatus int
		ExpectedRespBody string
	}{
		{
			Method: http.MethodGet,
			Endpoint: "/readiness",
			ExpectedRespStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Method + " " + tc.Endpoint, func(t *testing.T) {
			ctx, done := context.WithCancel(context.Background())
			t.Cleanup(func() {
				done()
			})

			listener, err := net.Listen("tcp", ":0")
			require.NoError(t, err)
			hostAddr := listener.Addr()

			go func() {
				err = runWithListener(ctx, listener)
				require.NoError(t, err)
			}()

			srvUrl, err := url.Parse("http://" + hostAddr.String())
			require.NoError(t, err)

			req, err := http.NewRequestWithContext(
				ctx,
				tc.Method,
				srvUrl.JoinPath(
					tc.Endpoint,
				).String(),
				strings.NewReader(tc.ReqBody),
			)
			require.NoError(t, err)

			client := http.DefaultClient

			resp, err := client.Do(req)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedRespStatus, resp.StatusCode)


			defer resp.Body.Close()
			respBody, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedRespBody, string(respBody))
		})
	}
}

func TestRun_DuplicatedEndpoints(t *testing.T) {

	endpoints := []string{
		"/aaa",
		"/aaa/aaa",
		"/aaa/bbb",
		"/aaa/ccc",
	}

	testCases := []struct{
		Method string
		Endpoint string
		ReqBody string
		ExpectedRespStatus int
		ExpectedRespBody string
	}{
		{
			Method: http.MethodGet,
			ReqBody: `{}`,
			ExpectedRespStatus: http.StatusOK,
			ExpectedRespBody: `{"items":[{"one_val":"one","two_val":1},{"one_val":"two","two_val":2}]}
`,
		},
		{
			Method: http.MethodPost,
			ReqBody: `{}`,
			ExpectedRespStatus: http.StatusOK,
			ExpectedRespBody: `{"success":true}
`,
		},
	}

	for _, endpoint := range endpoints {
		for _, tc := range testCases {
			t.Run(tc.Method + endpoint, func(t *testing.T) {
				ctx, done := context.WithCancel(context.Background())
				t.Cleanup(func() {
					done()
				})

				listener, err := net.Listen("tcp", ":0")
				require.NoError(t, err)
				hostAddr := listener.Addr()

				go func() {
					err = runWithListener(ctx, listener)
					require.NoError(t, err)
				}()

				srvUrl, err := url.Parse("http://" + hostAddr.String())
				require.NoError(t, err)

				req, err := http.NewRequestWithContext(
					ctx,
					tc.Method,
					srvUrl.JoinPath(
						endpoint,
					).String(),
					strings.NewReader(tc.ReqBody),
				)
				require.NoError(t, err)

				client := http.DefaultClient

				resp, err := client.Do(req)
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedRespStatus, resp.StatusCode)


				defer resp.Body.Close()
				respBody, err := ioutil.ReadAll(resp.Body)
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedRespBody, string(respBody))
			})
		}
	}
}

func TestRun_Expected404s(t *testing.T) {

	endpoints := []string{
		"/bbb",
		"/bbb/aaa",
		"/bbb/bbb",
		"/bbb/ccc",
		"/ccc",
		"/ccc/aaa",
		"/ccc/bbb",
		"/ccc/ccc",
	}

	testCases := []struct{
		Method string
		Endpoint string
		ReqBody string
		ExpectedRespStatus int
	}{
		{
			Method: http.MethodGet,
			ReqBody: `{}`,
			ExpectedRespStatus: http.StatusNotFound,
		},
		{
			Method: http.MethodPost,
			ReqBody: `{}`,
			ExpectedRespStatus: http.StatusNotFound,
		},
	}

	for _, endpoint := range endpoints {
		for _, tc := range testCases {
			t.Run(tc.Method + endpoint, func(t *testing.T) {
				ctx, done := context.WithCancel(context.Background())
				t.Cleanup(func() {
					done()
				})

				listener, err := net.Listen("tcp", ":0")
				require.NoError(t, err)
				hostAddr := listener.Addr()

				go func() {
					err = runWithListener(ctx, listener)
					require.NoError(t, err)
				}()

				srvUrl, err := url.Parse("http://" + hostAddr.String())
				require.NoError(t, err)

				req, err := http.NewRequestWithContext(
					ctx,
					tc.Method,
					srvUrl.JoinPath(
						endpoint,
					).String(),
					strings.NewReader(tc.ReqBody),
				)
				require.NoError(t, err)

				client := http.DefaultClient

				resp, err := client.Do(req)
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedRespStatus, resp.StatusCode)
			})
		}
	}
}
