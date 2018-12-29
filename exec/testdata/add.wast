(module
 (type $FUNCSIG$iii (func (param i32 i32) (result i32)))
 (import "env" "add" (func $add (param i32 i32) (result i32)))
 (table 0 anyfunc)
 (memory $0 1)
 (data (i32.const 4) "\10@\00\00")
 (export "memory" (memory $0))
 (export "add2" (func $add2))
 (export "main" (func $main))
 (func $add2 (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (i32.store offset=12
   (tee_local $2
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
   (get_local $2)
   (get_local $1)
  )
  (i32.add
   (i32.load offset=12
    (get_local $2)
   )
   (get_local $1)
  )
 )
 (func $main (result i32)
  (local $0 i32)
  (local $1 i32)
  (local $2 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $2
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (i32.store offset=12
   (get_local $2)
   (call $add
    (i32.const 1)
    (i32.const 2)
   )
  )
  (i32.store offset=8
   (get_local $2)
   (tee_local $0
    (call $add2
     (i32.const 3)
     (i32.const 4)
    )
   )
  )
  (set_local $1
   (i32.load offset=12
    (get_local $2)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $2)
    (i32.const 16)
   )
  )
  (i32.add
   (get_local $1)
   (get_local $0)
  )
 )
)
