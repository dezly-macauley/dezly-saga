; SECTION: Comments

; This is how to make a comment in Assembly

;______________________________________________________________________________
; SECTION: How to run this code


; nasm -f elf64 hello_world.asm -o hello_world.o && ld hello_world.o -o hello_world && ./hello_world
; The output will be this:
; Hello, World!

; nasm -f elf64 hello_world.asm -o hello_world.o 
; ld hello_world.o -o hello_world
; ./hello_world 

;______________________________________________________________________________
section .data
    msg db 'Hello, World!',0xa
    len equ $ - msg

section .text
    global _start

_start:
    mov rax, 1
    mov rdi, 1
    mov rsi, msg
    mov rdx, len
    syscall

    mov rax, 60
    xor rdi, rdi
    syscall
