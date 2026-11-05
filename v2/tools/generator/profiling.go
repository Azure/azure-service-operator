/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

type profiler struct {
	cpuProfilePath      string
	memoryProfilePath   string
	blockingProfilePath string
	cpuProfile          *os.File
	memoryProfile       *os.File
	blockingProfile     *os.File
}

func newProfiler() *profiler {
	return &profiler{}
}

func (p *profiler) start() error {
	var err error

	if p.memoryProfilePath != "" {
		p.memoryProfile, err = os.Create(p.memoryProfilePath)
		if err != nil {
			return fmt.Errorf("creating memory profile %q: %w", p.memoryProfilePath, err)
		}
	}

	if p.cpuProfilePath != "" {
		p.cpuProfile, err = os.Create(p.cpuProfilePath)
		if err != nil {
			return p.cleanupAfterStartFailure(
				fmt.Errorf("creating CPU profile %q: %w", p.cpuProfilePath, err),
			)
		}

		if err = pprof.StartCPUProfile(p.cpuProfile); err != nil {
			return p.cleanupAfterStartFailure(
				fmt.Errorf("starting CPU profile %q: %w", p.cpuProfilePath, err),
			)
		}
	}

	if p.blockingProfilePath != "" {
		p.blockingProfile, err = os.Create(p.blockingProfilePath)
		if err != nil {
			return p.cleanupAfterStartFailure(
				fmt.Errorf("creating blocking profile %q: %w", p.blockingProfilePath, err),
			)
		}

		runtime.SetBlockProfileRate(1)
	}

	return nil
}

func (p *profiler) stop() error {
	var result error

	if p.cpuProfile != nil {
		pprof.StopCPUProfile()
		result = errors.Join(result, p.closeCPUProfile())
	}

	if p.memoryProfile != nil {
		runtime.GC()
		allocations := pprof.Lookup("allocs")
		if allocations == nil {
			result = errors.Join(result, errors.New("finding cumulative allocation profile"))
		} else if err := allocations.WriteTo(p.memoryProfile, 0); err != nil {
			result = errors.Join(result, fmt.Errorf("writing memory profile %q: %w", p.memoryProfilePath, err))
		}

		result = errors.Join(result, p.closeMemoryProfile())
	}

	if p.blockingProfile != nil {
		runtime.SetBlockProfileRate(0)
		profile := pprof.Lookup("block")
		if profile != nil {
			err := profile.WriteTo(p.blockingProfile, 0)
			if err != nil {
				result = errors.Join(result, fmt.Errorf("writing blocking profile %q: %w", p.blockingProfilePath, err))
			}
		}

		result = errors.Join(result, p.closeBlockingProfile())
	}

	return result
}

func (p *profiler) cleanupAfterStartFailure(startErr error) error {
	return errors.Join(
		startErr,
		p.discardCPUProfile(),
		p.discardMemoryProfile(),
		p.discardBlockingProfile(),
	)
}

func (p *profiler) closeCPUProfile() error {
	if p.cpuProfile == nil {
		return nil
	}

	err := p.cpuProfile.Close()
	p.cpuProfile = nil
	if err != nil {
		return fmt.Errorf("closing CPU profile %q: %w", p.cpuProfilePath, err)
	}

	return nil
}

func (p *profiler) discardCPUProfile() error {
	if p.cpuProfile == nil {
		return nil
	}

	return errors.Join(
		p.closeCPUProfile(),
		p.removeCPUProfile(),
	)
}

func (p *profiler) closeMemoryProfile() error {
	if p.memoryProfile == nil {
		return nil
	}

	err := p.memoryProfile.Close()
	p.memoryProfile = nil
	if err != nil {
		return fmt.Errorf("closing memory profile %q: %w", p.memoryProfilePath, err)
	}

	return nil
}

func (p *profiler) discardMemoryProfile() error {
	if p.memoryProfile == nil {
		return nil
	}

	return errors.Join(
		p.closeMemoryProfile(),
		p.removeMemoryProfile(),
	)
}

func (p *profiler) removeCPUProfile() error {
	if err := os.Remove(p.cpuProfilePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("removing CPU profile %q: %w", p.cpuProfilePath, err)
	}

	return nil
}

func (p *profiler) removeMemoryProfile() error {
	if err := os.Remove(p.memoryProfilePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("removing memory profile %q: %w", p.memoryProfilePath, err)
	}

	return nil
}

func (p *profiler) closeBlockingProfile() error {
	if p.blockingProfile == nil {
		return nil
	}

	err := p.blockingProfile.Close()
	p.blockingProfile = nil
	if err != nil {
		return fmt.Errorf("closing blocking profile %q: %w", p.blockingProfilePath, err)
	}

	return nil
}

func (p *profiler) discardBlockingProfile() error {
	if p.blockingProfile == nil {
		return nil
	}

	return errors.Join(
		p.closeBlockingProfile(),
		p.removeBlockingProfile(),
	)
}

func (p *profiler) removeBlockingProfile() error {
	if err := os.Remove(p.blockingProfilePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("removing blocking profile %q: %w", p.blockingProfilePath, err)
	}

	return nil
}
