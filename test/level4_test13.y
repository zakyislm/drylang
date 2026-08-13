
    try {
        throw_PKzOFb = unknown
        throw_PKzOFb()
    } err(e) {
        pt("caught error")
    }
    

    fn calc_cDaKiF(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_hgRkDp = calc_cDaKiF(42, 88+1)
    pt(res_hgRkDp)
    

    val_UspgaR = 42 % 3
    on (val_UspgaR) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_UspgaR)
    

    cl Base_gFOqbY {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_Qlyhgv <- Base_gFOqbY {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_tKVKXM = Child_Qlyhgv(26, 21)
    pt(obj_tKVKXM.get_id())
    

    use "dummy.y"
    
    e_MvrTcA = enc.b64("hello drylang")
    pt(e_MvrTcA)
    
    j_bplwMR = json(`{"test": 123}`)
    pt(j_bplwMR)
    
    // Test get() for type info
    val_khzyYh = get("hello")
    pt(val_khzyYh)
    

    arr_lyqkUa = [69, 65+1, 100+2]
    map_rPqjUQ = {"a": arr_lyqkUa[0], "b": arr_lyqkUa[1]}
    pt(map_rPqjUQ["a"])
    
