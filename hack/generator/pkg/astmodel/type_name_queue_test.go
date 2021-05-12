/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMakeTypeNameQueue_Len_IsZero(t *testing.T) {
	g := NewGomegaWithT(t)

	queue := MakeTypeNameQueue()
	g.Expect(queue.Len()).To(Equal(0))
}

func TestTypeNameQueue_Enqueue(t *testing.T) {
	g := NewGomegaWithT(t)
	email := MakeTypeName(emailTestRef, "Email")

	queue := MakeTypeNameQueue()
	initialLength := queue.Len()

	queue.Enqueue(email)

	g.Expect(queue.Len() > initialLength).To(BeTrue())
}

func TestTypeNameQueue_Dequeue(t *testing.T) {
	g := NewGomegaWithT(t)
	email := MakeTypeName(emailTestRef, "Email")

	queue := MakeTypeNameQueue()
	queue.Enqueue(email)

	initialLength := queue.Len()

	n, ok := queue.Dequeue()

	g.Expect(ok).To(BeTrue())
	g.Expect(queue.Len() < initialLength).To(BeTrue())
	g.Expect(n).To(Equal(email))

	_, ok = queue.Dequeue()
	g.Expect(ok).To(BeFalse())
}

func TestTypeNameQueue_Process(t *testing.T) {
	g := NewGomegaWithT(t)
	email := MakeTypeName(emailTestRef, "Email")
	inbox := MakeTypeName(emailTestRef, "Inbox")
	mailingList := MakeTypeName(emailTestRef, "MailingList")

	queue := MakeTypeNameQueue()
	queue.Enqueue(email)
	queue.Enqueue(inbox)
	queue.Enqueue(mailingList)

	initialQueueSize := queue.Len()
	processed := 0

	err := queue.Process(func(n TypeName) error {
		processed++
		return nil
	})

	g.Expect(err).To(BeNil())
	g.Expect(processed).To(Equal(initialQueueSize))
}
