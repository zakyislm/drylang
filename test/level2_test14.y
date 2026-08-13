
    cl Base_Whyulf {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_FwjRyl <- Base_Whyulf {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_iFNIUb = Child_FwjRyl(44, 48)
    pt(obj_iFNIUb.get_id())
    

    val_LfbXWR = 55 % 3
    on (val_LfbXWR) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_LfbXWR)
    

    try {
        throw_CZluTL = unknown
        throw_CZluTL()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_bHQtRo = enc.b64("hello drylang")
    pt(e_bHQtRo)
    
    j_DUAmFu = json(`{"test": 123}`)
    pt(j_DUAmFu)
    
    // Test get() for type info
    val_CMGsGO = get("hello")
    pt(val_CMGsGO)
    
