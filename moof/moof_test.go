package moof

import "testing"

func TestRead

/*
def read(self, data) :
    cursor = data.cursor
    size = data.read(4)
    name = data.read(4)
    assert name == self.BOXTYPE
    self.boxes = [ ]
    while data.cursor - cursor < size :
      self.boxes.append(Box().read(data))
    return self
*/

