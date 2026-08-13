
    asn fn async_task_YHRopL(x) {
        rev x * 2
    }
    uni async_task_YHRopL(52)
    
    asn fn worker_twJHFM(y) {
        pt("working on", y)
    }
    mul 2 worker_twJHFM(80)
    
    awt
    

    m_BgkotB = math.sqrt(88)
    h_joJeWB = hash.md5("test_WXTodD")
    log.inf("hash:", h_joJeWB)
    pt(m_BgkotB)
    

    use "dummy.y"
    
    e_dclbhO = enc.b64("hello drylang")
    pt(e_dclbhO)
    
    j_grJyna = json(`{"test": 123}`)
    pt(j_grJyna)
    
    // Test get() for type info
    val_jRRKVu = get("hello")
    pt(val_jRRKVu)
    

    cl Base_ojPAVD {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_FqqkRe <- Base_ojPAVD {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_vPzDgP = Child_FqqkRe(63, 19)
    pt(obj_vPzDgP.get_id())
    

    cl Base_utJQhq {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_eeuEOk <- Base_utJQhq {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_mVewfk = Child_eeuEOk(68, 94)
    pt(obj_mVewfk.get_id())
    

    try {
        throw_EAMCpc = unknown
        throw_EAMCpc()
    } err(e) {
        pt("caught error")
    }
    
