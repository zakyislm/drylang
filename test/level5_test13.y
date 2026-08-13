
    arr_TggCcW = [71, 11+1, 43+2]
    map_QGPoYR = {"a": arr_TggCcW[0], "b": arr_TggCcW[1]}
    pt(map_QGPoYR["a"])
    

    arr_NeNjFY = [26, 11+1, 19+2]
    map_bsDVVZ = {"a": arr_NeNjFY[0], "b": arr_NeNjFY[1]}
    pt(map_bsDVVZ["a"])
    

    cl Base_QzhMAk {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_CiZHNv <- Base_QzhMAk {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_lspmDR = Child_CiZHNv(31, 68)
    pt(obj_lspmDR.get_id())
    

    val_zcVRBx = 2 % 3
    on (val_zcVRBx) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_zcVRBx)
    

    arr_qbvESD = [6, 23+1, 55+2]
    map_piQKZN = {"a": arr_qbvESD[0], "b": arr_qbvESD[1]}
    pt(map_piQKZN["a"])
    

    val_ckWoaG = 42 % 3
    on (val_ckWoaG) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_ckWoaG)
    

    try {
        throw_uJEOCp = unknown
        throw_uJEOCp()
    } err(e) {
        pt("caught error")
    }
    
