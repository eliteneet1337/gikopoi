/**
 * Utility functions for the project.
 */
package main

import (
  "github.com/google/uuid"
)


/**
 * Generate a new UUID string and return it.
 */
func makeUuid() string {
  return uuid.New().String()
}
