package main

import (
	"fmt"
	"strings"
	"testing"
)

var largeString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. " + "Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. " + "Praesent et diam eget libero egestas mattis sit amet vitae augue. Nam tincidunt congue enim, " + "ut porta lorem lacinia consectetur. Donec ut libero sed arcu vehicula ultricies a non tortor. " + "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean ut gravida lorem. " + "Ut turpis felis, pulvinar a semper sed, adipiscing id dolor. Pellentesque auctor nisi id magna consequat " + "sagittis. Curabitur dapibus enim sit amet elit pharetra tincidunt feugiat nisl imperdiet. " + "Ut convallis libero in urna ultrices accumsan. Donec sed odio eros. Donec viverra mi quis quam " + "pulvinar at malesuada arcu rhoncus. Cum sociis natoque penatibus et magnis dis parturient montes, " + "nascetur ridiculus mus. In rutrum accumsan ultricies. Mauris vitae nisi at sem facilisis semper ac " + "in est." + "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.Lorem ipsum dolor sit amet, qui minim labore adipisicing hello minim sint cillum sint consectetur cupidatat."

func init() {
	fmt.Println(len(largeString))

	fmt.Println(len(largeString))
}

func BenchmarkCut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, _, ok := strings.Cut(largeString, "hello"); !ok {
			continue
		}
	}
}

func BenchmarkLastIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := strings.LastIndex(largeString, "hello")
		if i == -1 {
			continue
		}
		b := largeString[:i]
		_ = b
	}
}
