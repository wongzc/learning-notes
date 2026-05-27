1. what
    - X86 reset state
    - X86 use segment:offset addressing in real mode
        - CS: code segment, base location
        - IP: instruction pointer, offset in segment
    - when CPU power on, start execute at CS=0xF000 and IP=0xFFF0
        - physical= CS*16+IP= 0xFFFF0 (16byte below/before 1MB (0x100000, 2^20))
            - *16 == <<4
            -  x86 (8086) register only 16 bit, max 2^16 bytes= 65536 bytes = 64 KB
            - so segment <<4 + offset, means segment rep a blocks every 16 byte, can rep 20 bit address

2. why limit?
    - tiny startup space
        - start at 0xFFFF0, only 16byte before 1MB, almost no room
            - total firmware is from 0xF0000 to 0xFFFFF, so many space
            - but entry is only left with 16 byte ( 0xFFFF0 to 0xFFFFF)
        - CPU cant start directly in larger OS kernel
            - only tiny stub here
        - only enough for `jmp far actual_bios_start`: to jump to bios start
    - force dependency on BIOS firmware
        - cannot start kernel directly, boot more complex and slow
            - CPU reset vector hardcoded to firmware address at 0xFFFF0
            - not at kernel
        - flow:
            1. CPU jumps to BIOS ROM ( run firmware)
            2. BIOS run POST
            3. BIOS find boot device
            4. BIOS load bootloader
            5. bootloader loads OS
    - read mode limitation
        - CPU start in 16 bit real mode, constraint
            - 1MB address space
            - segmented memory
            - no paging, memory protect, privilege level, virtual memory
        - need to go from real mode -> protected mode -> long mode  
    - 512-byte boot sector limitation
        - BIOS boot load first sector (512 byte)
        - first-stage bootloader need to fit in 512 byte
        - so need multiple stage load
    - legacy compatibility baggage
        - x86 with many compatibility hacks:
            1. A20 line weirdness
            2. BIOS interrupts
            3. real mode shims
            4. protected mode switching
            5. AP startup tricks in SMP

3. ROM BIOS
    - ROM: read only memory
        - old PC firmware stored in ROM chip in motherboard
            - firmware: code that exists before OS
                - initilaize hardware
                - POST (power-on self test)
                - detect disk/ keyboard/ video
                - boot OS
    - BIOS: Basic Input/Output System
    - modern day use UEFI
    - BIOS ROM: from 0xF0000 to 0xFFFFF ( 64KB)

4. X86:
    - ISA: instruction set architecture
        - define machine instructions, registers, memory model, privilege modes, interrupt model, calling conventions, boot/ reset behavior
        - define how software communicate with CPU
        - for CPU only ( not GPU, SSD etc)
    - mainstream today for PC/ server
        - Intel/AMD
    - phone is ARM ( iphone/ android)
        - apple M chip, AWS graviton also ARM
    - x86 means chip ending with 86
        - 8086, 80186, 80286 ...
        - modern means: x86-64 ( 64 bit extension)

5. Summary:
    - true issue
        - CPU reset hardcoded to start in firmware, not OS/bootloader
            - CPU reset: CPU put into known initial state
        - need to do transition from real -> protected -> maybe long
        - bootloader need multiple stage because of 512 byte
          - beacuse BIOS read 1 disk sector (512 byte) from boot device 
          - BIOS boot protocol load firect sector and execute it
            - why only 1? historical reason, so just read sectr 0 and jump
          - first stage must fit in 512 byte 
        - make X86 boot and early intialization complex
    - flow now:
        1. power on, CPU reset
        2. CS:Ip = F000:FFF0, physical=0xFFFF0
        3. execute firmware stub ( just jump only)
        4. jump to BIOS firware code
        5. BIOS init hardware
        6. BIOS load bootloader
        7. bootloader load OS
    - if reset not hardcoded
        1. power on, CPU reset
        2. CPU start at confogurable boot entry
        3. kernel start
        4. init hardware
        5. run OS


