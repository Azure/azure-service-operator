/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// TypeNameQueue is a classic FIFO queue of TypeNames
type TypeNameQueue struct {
	queue []TypeName
}

// MakeTypeNameQueue creates a new empty queue
func MakeTypeNameQueue() TypeNameQueue {
	return TypeNameQueue{}
}

// Enqueue adds the specified TypeName to back of the queue
// Duplicates are allowed
func (q *TypeNameQueue) Enqueue(name TypeName) {
	q.queue = append(q.queue, name)
}

// Dequeue removes the front item from the queue
// An error will be returned if no item is available
func (q *TypeNameQueue) Dequeue() (TypeName, bool) {
	if len(q.queue) == 0 {
		return TypeName{}, false
	}

	// TODO: There's a gotcha here that the slices retain the underlying array
	// so along lived queue that never empties will prevent a bunch of stuff
	// from being cleaned up by the GC
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result, true
}

// Len returns the length of the queue
func (q *TypeNameQueue) Len() int {
	return len(q.queue)
}

// Process will drain the queue, passing each item in turn to the supplied func
// It's permissible for the func to add more items to the queue for processing
// If the func returns an error, processing of the queue will terminate at that point, with the
// error returned
func (q *TypeNameQueue) Process(processor func(TypeName) error) error {
	for len(q.queue) > 0 {
		def, ok := q.Dequeue()
		if !ok {
			// Emptied the queue, we're done
			return nil
		}

		err := processor(def)
		if err != nil {
			return err
		}
	}

	return nil
}
