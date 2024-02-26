package main

import (
	"context"
	"log"
)

// Tool that mocks a weather forecast
type WeatherForecastMock struct{}

func NewWeatherForecastMock() *WeatherForecastMock {
	return &WeatherForecastMock{}
}

func (magicNumber *WeatherForecastMock) Name() string {
	return "weather"
}

func (magicNumber *WeatherForecastMock) Description() string {
	return `Took that allows to retrieve the live weather forecast for a location.
Always provide the input as "City, Province/State, Country"`
}

func (magicNumber *WeatherForecastMock) Call(ctx context.Context, input string) (string, error) {
	log.Printf("WeatherForecastMock::input = %q", input)
	return "25 and Sunny", nil
}
