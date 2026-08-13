
    val_SSmnqU = 15 % 3
    on (val_SSmnqU) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_SSmnqU)
    

    sum_Ckkxek = 0
    lp 12 {
        sum_Ckkxek = sum_Ckkxek + 1
        if (sum_Ckkxek > 100) { done }
    }
    pt(sum_Ckkxek)
    

    m_UUTtRw = math.sqrt(18)
    h_cQnVuK = hash.md5("test_ceXIRU")
    log.inf("hash:", h_cQnVuK)
    pt(m_UUTtRw)
    

    use "dummy.y"
    
    e_iQpSYj = enc.b64("hello drylang")
    pt(e_iQpSYj)
    
    j_NkewdW = json(`{"test": 123}`)
    pt(j_NkewdW)
    
    // Test get() for type info
    val_KcjBxJ = get("hello")
    pt(val_KcjBxJ)
    

    use "dummy.y"
    
    e_pDfiCY = enc.b64("hello drylang")
    pt(e_pDfiCY)
    
    j_fDJXfm = json(`{"test": 123}`)
    pt(j_fDJXfm)
    
    // Test get() for type info
    val_dSIYZy = get("hello")
    pt(val_dSIYZy)
    

    cl Base_sYFdDJ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_XayhVU <- Base_sYFdDJ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_IjrozG = Child_XayhVU(88, 49)
    pt(obj_IjrozG.get_id())
    
