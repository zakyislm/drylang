
    val_Uqdbrw = 51 % 3
    on (val_Uqdbrw) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_Uqdbrw)
    

    sum_Yvjtzd = 0
    lp 15 {
        sum_Yvjtzd = sum_Yvjtzd + 1
        if (sum_Yvjtzd > 100) { done }
    }
    pt(sum_Yvjtzd)
    

    try {
        throw_FttLqd = unknown
        throw_FttLqd()
    } err(e) {
        pt("caught error")
    }
    

    sum_gJmlGt = 0
    lp 9 {
        sum_gJmlGt = sum_gJmlGt + 1
        if (sum_gJmlGt > 100) { done }
    }
    pt(sum_gJmlGt)
    

    use "dummy.y"
    
    e_WvDNaT = enc.b64("hello drylang")
    pt(e_WvDNaT)
    
    j_arhKlw = json(`{"test": 123}`)
    pt(j_arhKlw)
    
    // Test get() for type info
    val_uTKIns = get("hello")
    pt(val_uTKIns)
    

    use "dummy.y"
    
    e_HveHER = enc.b64("hello drylang")
    pt(e_HveHER)
    
    j_AAkQTC = json(`{"test": 123}`)
    pt(j_AAkQTC)
    
    // Test get() for type info
    val_VrOWMN = get("hello")
    pt(val_VrOWMN)
    

    val_msNzjj = 75 % 3
    on (val_msNzjj) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_msNzjj)
    

    cl Base_IOzKTj {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_gKMFOA <- Base_IOzKTj {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_GqsGHA = Child_gKMFOA(89, 34)
    pt(obj_GqsGHA.get_id())
    
