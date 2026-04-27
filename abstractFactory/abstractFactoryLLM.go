package main

import (
    "fmt"
    "interfaces"
)



func GetSportsFactory(brand string) (interfaces.IPromptFactory, error) {
    if brand == "adidas" {
        return &{}, nil
    }

    return nil, fmt.Errorf("Wrong brand type passed")
}