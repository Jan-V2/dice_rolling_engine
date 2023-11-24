package main

import "fmt"

type Vertex_Attribute struct{
    name string
    num_floats int
}

type Vertex_Buffer struct{
    buffer []float32
    indexes []uint32
    attributes []Vertex_Attribute
    floats_per_vertex int
}

func new_vertex_buffer()Vertex_Buffer{
    return Vertex_Buffer{
        buffer:            nil,
        attributes:        nil,
        floats_per_vertex: 0,
    }
}

func (v *Vertex_Buffer) add_attribute(attr Vertex_Attribute){
    v.floats_per_vertex += attr.num_floats
    v.attributes = append(v.attributes, attr)
}

func (v *Vertex_Buffer) add_vertex(points []float32)  {
    if len(points) != v.floats_per_vertex{
        err := fmt.Sprintf("incorrect number of points in vertex, was expecting %d got %d",
            v.floats_per_vertex, len(points))
        panic(err)
    }
    v.buffer = append(v.buffer, points...)
}

func (v *Vertex_Buffer) get_stride() int32 {
    var ret int32
    for _, attribute := range v.attributes {
        ret += int32(attribute.get_size())
    }
    return ret
}



func (v *Vertex_Buffer) get_len() int32 {
    return int32(len(v.buffer) / v.floats_per_vertex)
}
func (v *Vertex_Attribute) get_size() int{
    return SIZEOF_FLOAT32 * v.num_floats
}

