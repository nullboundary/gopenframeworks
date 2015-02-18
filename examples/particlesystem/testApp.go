package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	gof "github.com/nullboundary/gopenframeworks"
	"math/rand"
	"time"
)

type particleSystem struct {
	particles []*particle //list of particles
	origin    mgl32.Vec3  //center or start of the particle system
}

type particle struct {
	loc      mgl32.Vec3 //location
	vel      mgl32.Vec3 //velocity
	acc      mgl32.Vec3 //acceleration
	lifespan float32
}

var width float32 = 1024
var height float32 = 768
var ps *particleSystem
var r = rand.New(rand.NewSource(time.Now().UnixNano())) //random seed
var totalParticles = 100                                //desired number of particles in the system

// setup
//--------------------------------------------------------------
func (app testApp) Setup() {
	fmt.Println("Particle System")
	gof.Background(0, 0, 0.2)
	ps = newParticleSystem(mgl32.Vec3{20, 50, 0})

}

// update
//--------------------------------------------------------------
func (app testApp) Update() {
	ps.update()
	ps.addParticle()
}

// Draw in the window
//--------------------------------------------------------------
func (app testApp) Draw() {
	ps.draw()
}

// create a new particleSystem
//--------------------------------------------------------------
func newParticleSystem(location mgl32.Vec3) *particleSystem {
	ps := new(particleSystem)
	ps.origin = location
	ps.particles = make([]*particle, 0, 512)
	return ps
}

// add a number of particles at once
//--------------------------------------------------------------
func (ps *particleSystem) build(numParticles int) {
	for i := 0; i <= numParticles; i++ {
		fmt.Println(i)
		ps.addParticle()
	}
}

// add a particle to the system
//--------------------------------------------------------------
func (ps *particleSystem) addParticle() {
	p := newParticle(ps.origin)
	ps.particles = append(ps.particles, p)
}

// update the values of each particle in the particle System
//--------------------------------------------------------------
func (ps *particleSystem) update() {
	fmt.Println(len(ps.particles))
	for _, p := range ps.particles {
		p.update()
	}

	if len(ps.particles) >= totalParticles {
		//pop front
		ps.particles = ps.particles[:len(ps.particles)-1]
	}

}

// Draw the whole particle system
//--------------------------------------------------------------
func (ps *particleSystem) draw() {
	for i, _ := range ps.particles {
		p := ps.particles[i]
		p.draw()
	}
}

// Creates a particle
//--------------------------------------------------------------
func newParticle(l mgl32.Vec3) *particle {
	p := new(particle)
	p.acc = mgl32.Vec3{r.Float32() / 10, r.Float32() / 10, 0}
	p.vel = mgl32.Vec3{r.Float32(), r.Float32(), 0}
	p.loc = l
	p.lifespan = 255.0

	return p
}

// update the particle vel,loc,lifespan
//--------------------------------------------------------------
func (p *particle) update() {
	p.vel = p.vel.Add(p.acc)
	p.loc = p.loc.Add(p.vel)
	p.lifespan = p.lifespan - 2.0
}

// draw it
//--------------------------------------------------------------
func (p *particle) draw() {
	partShape := gof.Rectangle(p.loc[0], p.loc[1], 8, 8) //x,y,w,h
	partShape.Fill(0, 1, 0.5, 0.2)                       //r,g,b,a
	partShape.Draw()
}

//--------------------------------------------------------------
func (app testApp) KeyPressed(key int) {}

//--------------------------------------------------------------
func (app testApp) KeyReleased(key int) {}

//--------------------------------------------------------------
func (app testApp) MouseMoved(x int, y int) {}

//--------------------------------------------------------------
func (app testApp) MouseDragged(x int, y int, button int) {}

//--------------------------------------------------------------
func (app testApp) MousePressed(x int, y int, button int) {}

//--------------------------------------------------------------
func (app testApp) MouseReleased(x int, y int, button int) {}

//--------------------------------------------------------------
func (app testApp) WindowResized(w int, h int) {}
