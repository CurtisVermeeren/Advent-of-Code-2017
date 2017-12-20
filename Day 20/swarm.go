package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type coord struct {
	x, y, z int
}

type particle struct {
	position, acceleration, velocity coord
	collided                         bool
}

func main() {
	particles := getInput()
	part1 := solvePart1(particles)
	fmt.Println("Part 1:", part1)
	parts2 := solvePart2(particles)
	fmt.Println("Part 2:", parts2)
}

// Function builds a slice of particle objects from the file of particles
func getInput() []particle {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	particles := make([]particle, 0)

	for scanner.Scan() {
		line := scanner.Text()
		// Read the line, build the particle and get the values using Sscanf
		p := &particle{}
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.position.x, &p.position.y, &p.position.z,
			&p.velocity.x, &p.velocity.y, &p.velocity.z,
			&p.acceleration.x, &p.acceleration.y, &p.acceleration.z)
		particles = append(particles, *p)
	}
	return particles
}

// Function updates a particles position for that tick
func updateParticle(p *particle) {
	// Update the velocity using acceleration
	p.velocity.x += p.acceleration.x
	p.velocity.y += p.acceleration.y
	p.velocity.z += p.acceleration.z
	// Update the position using velocity
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.position.z += p.velocity.z
}

// Function solves part 1
func solvePart1(particles []particle) int {
	// Track the minimum index and minimum value.
	minIndex := 0
	aMin := ^uint(0)
	// For each particle we calculate the absolute of the acceleration.
	// The longer a particle moves the further it will eventually become to (0,0,0)
	// We want to find the particle with the smallest amount of acceleration because that will be the one that moves the least and stays closest to 0,0,0 longest.
	for index, particle := range particles {
		a := math.Abs(float64(particle.acceleration.x)) + math.Abs(float64(particle.acceleration.y)) + math.Abs(float64(particle.acceleration.z))
		if uint(a) <= aMin {
			aMin = uint(a)
			minIndex = index
		}
	}
	return minIndex
}

func solvePart2(particles []particle) int {
	length := len(particles)

	oldLength := length
	moving := true
	for i := 1; moving; i++ {
		// Track all positions of particles at the current tick
		positions := make(map[coord]*particle)
		for index := range particles {
			p := &particles[index]

			// If the particle has collided skip it
			if p.collided {
				continue
			}

			// If there is not a particle at this position then add this particle at this positon
			if _, ok := positions[p.position]; !ok {
				positions[p.position] = p
				// If there is a particle at the position the destroy them both
			} else {
				// Set the particle as collided and decrase the total number of particles
				p.collided = true
				length--
				if other := positions[p.position]; !other.collided {
					other.collided = true
					length--
				}

			}
			// Update the particle's position
			updateParticle(p)
		}
		// The particles will move forever unless we stop them
		// Every 100 movements we check if there have been more collisions than the last check
		if i%100 == 0 {
			// If there have been no new collisions in this time stop execution
			if length == oldLength {
				moving = false
			}
			oldLength = length
		}
	}
	return length
}
