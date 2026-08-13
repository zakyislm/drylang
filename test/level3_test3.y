
    try {
        throw_ZlhZTy = unknown
        throw_ZlhZTy()
    } err(e) {
        pt("caught error")
    }
    

    arr_evoQso = [11, 87+1, 21+2]
    map_svFbEm = {"a": arr_evoQso[0], "b": arr_evoQso[1]}
    pt(map_svFbEm["a"])
    

    use "dummy.y"
    
    e_gMSxlC = enc.b64("hello drylang")
    pt(e_gMSxlC)
    
    j_yCUXjF = json(`{"test": 123}`)
    pt(j_yCUXjF)
    
    // Test get() for type info
    val_hgnUav = get("hello")
    pt(val_hgnUav)
    

    cl Base_FYoPnA {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_lwvocV <- Base_FYoPnA {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ewejPn = Child_lwvocV(92, 31)
    pt(obj_ewejPn.get_id())
    

    fn calc_CeYdTP(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_WgeAVB = calc_CeYdTP(89, 55+1)
    pt(res_WgeAVB)
    
