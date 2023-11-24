package main

const (
    vertexShaderSource = `
        #version 410
        layout (location = 0) in vec3 aPos;
        layout (location = 1) in vec3 aColor;
        layout (location = 2) in vec2 aTexCoord;
        
        out vec3 frag_color;
        out vec2 tex_coord;

        void main() {
            gl_Position = vec4(aPos, 1.0);
            frag_color =  aColor;
            tex_coord = aTexCoord;
        }
    ` + "\x00"

    fragmentShaderSource = `
        #version 410
        in vec3 frag_color;
        in vec2 tex_coord;
        out vec4 color;

        uniform sampler2D tex;        

        void main() {
            //vec4(frag_color, 1.0);
            color = texture(tex, tex_coord);
        }
    ` + "\x00"
)
