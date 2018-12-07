package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestComputePart1(t *testing.T) {
	type testCase struct {
		name string
		in   string
		out  int
	}

	cases := []testCase{
		{
			name: "small",
			in:   "in1.txt",
			out:  0,
		},
		{
			name: "example",
			in:   "in2.txt",
			out:  17,
		},
		{
			name: "part1",
			in:   "in3.txt",
			out:  3871,
		},
		{
			name: "part1-more",
			in:   "in4.txt",
			out:  4284,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			coords, err := fileToCoords(tc.in)
			if err != nil {
				t.Fatal(err)
			}

			out := computePart1(coords)

			if tc.out != out.area {
				t.Errorf("expected %v, got %v", tc.out, out.area)
			}
		})
	}
}

func TestComputePart2(t *testing.T) {
	type testCase struct {
		name    string
		in      string
		maxDist int
		out     int
	}

	cases := []testCase{
		{
			name:    "example",
			in:      "in2.txt",
			maxDist: 32,
			out:     16,
		},
		{
			name:    "part2",
			in:      "in3.txt",
			maxDist: 10000,
			out:     44667,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			coords, err := fileToCoords(tc.in)
			if err != nil {
				t.Fatal(err)
			}

			out := computePart2(coords, tc.maxDist)

			if tc.out != out {
				t.Errorf("expected %v, got %v", tc.out, out)
			}
		})
	}
}

func fileToCoords(filename string) ([]*coord, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var coords []*coord
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), ", ")
		x, err := strconv.Atoi(points[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(points[1])
		if err != nil {
			return nil, err
		}
		coords = append(coords, &coord{
			x: x,
			y: y,
		})
	}

	return coords, scanner.Err()
}
