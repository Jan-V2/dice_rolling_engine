package main

const (
    vertexShaderSource = `
        #version 410
        layout (location = 0) in vec3 aPos;
        layout (location = 1) in vec2 aTexCoord;
        
        out vec2 tex_coord;
        uniform mat4 model;
        uniform mat4 view;
        uniform mat4 projection;
        
        void main() {
            gl_Position = projection * view * model * vec4(aPos, 1.0);
            tex_coord = aTexCoord;
        }
    ` + "\x00"

    fragmentShaderSource = `
        #version 410
        in vec2 tex_coord;
        out vec4 color;

        uniform sampler2D tex;        

        void main() {
            //vec4(frag_color, 1.0);
            color = texture(tex, tex_coord);
        }
    ` + "\x00"
)
