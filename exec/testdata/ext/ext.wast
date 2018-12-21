(module
 (table 0 anyfunc)
 (memory $0 1)
 (data (i32.const 4) "\10@\00\00")
 (export "memory" (memory $0))
 (export "main" (func $main))
 (export "grow" (func $grow))
 (export "callDepth2" (func $callDepth2))
 (export "callDepth" (func $callDepth))
 (func $main (param $0 i32) (result i32)
  (local $1 i32)
  (i32.store offset=12
   (tee_local $1
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
   (get_local $0)
  )
  (i32.store offset=8
   (get_local $1)
   (i32.const 0)
  )
  (i32.store offset=4
   (get_local $1)
   (i32.const 0)
  )
  (block $label$0
   (loop $label$1
    (br_if $label$0
     (i32.ge_s
      (i32.load offset=4
       (get_local $1)
      )
      (i32.load offset=12
       (get_local $1)
      )
     )
    )
    (i32.store offset=8
     (get_local $1)
     (i32.add
      (i32.load offset=8
       (get_local $1)
      )
      (i32.load offset=4
       (get_local $1)
      )
     )
    )
    (i32.store offset=4
     (get_local $1)
     (i32.add
      (i32.load offset=4
       (get_local $1)
      )
      (i32.const 1)
     )
    )
    (br $label$1)
   )
  )
  (i32.load offset=8
   (get_local $1)
  )
 )
 (func $grow (param $0 i32) (result i32)
  (local $1 i32)
  (i32.store offset=12
   (tee_local $1
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
   (get_local $0)
  )
  (i32.store offset=8
   (get_local $1)
   (i32.const 0)
  )
  (block $label$0
   (loop $label$1
    (br_if $label$0
     (i32.ge_s
      (i32.load offset=8
       (get_local $1)
      )
      (i32.load offset=12
       (get_local $1)
      )
     )
    )
    (drop
     (grow_memory
      (i32.const 1)
     )
    )
    (i32.store offset=8
     (get_local $1)
     (i32.add
      (i32.load offset=8
       (get_local $1)
      )
      (i32.const 1)
     )
    )
    (br $label$1)
   )
  )
  (i32.const 0)
 )
 (func $callDepth2 (param $0 i32) (result i32)
  (local $1 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $1
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (i32.store offset=8
   (get_local $1)
   (get_local $0)
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (get_local $0)
     )
    )
    (i32.store offset=12
     (get_local $1)
     (i32.add
      (tee_local $0
       (i32.load offset=8
        (get_local $1)
       )
      )
      (call $callDepth
       (i32.add
        (get_local $0)
        (i32.const -1)
       )
      )
     )
    )
    (br $label$0)
   )
   (i32.store offset=12
    (get_local $1)
    (i32.const 0)
   )
  )
  (set_local $0
   (i32.load offset=12
    (get_local $1)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $1)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $callDepth (param $0 i32) (result i32)
  (local $1 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $1
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (i32.store offset=8
   (get_local $1)
   (get_local $0)
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (get_local $0)
     )
    )
    (i32.store offset=12
     (get_local $1)
     (i32.add
      (tee_local $0
       (i32.load offset=8
        (get_local $1)
       )
      )
      (call $callDepth2
       (i32.add
        (get_local $0)
        (i32.const -1)
       )
      )
     )
    )
    (br $label$0)
   )
   (i32.store offset=12
    (get_local $1)
    (i32.const 0)
   )
  )
  (set_local $0
   (i32.load offset=12
    (get_local $1)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $1)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
)
