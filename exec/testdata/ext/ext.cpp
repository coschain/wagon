extern "C" {
    int main(int argc) {

        int value = 0;

        for ( int i = 0; i < argc; i++){
            value += i;
            value << argc;
        }
        return value;
    }

    int grow(int argc) {

        for ( int i = 0 ; i < argc; i++)
            __builtin_wasm_grow_memory(1);

        return 0;
    }

    int callDepth(int argc);
    int callDepth2(int argc) {

        if ( argc == 0 )
           return 0;

        return argc + callDepth(argc-1);
    }

    int callDepth(int argc) {

        if ( argc == 0 )
           return 0;

        return argc + callDepth2(argc-1);
    }
}