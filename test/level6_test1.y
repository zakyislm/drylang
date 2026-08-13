
    fn calc_KGQZLI(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_VvihVE = calc_KGQZLI(28, 56+1)
    pt(res_VvihVE)
    

    asn fn async_task_XodKlH(x) {
        rev x * 2
    }
    uni async_task_XodKlH(96)
    
    asn fn worker_KRzJzV(y) {
        pt("working on", y)
    }
    mul 2 worker_KRzJzV(4)
    
    awt
    

    use "dummy.y"
    
    e_UdJnqx = enc.b64("hello drylang")
    pt(e_UdJnqx)
    
    j_lZBBfH = json(`{"test": 123}`)
    pt(j_lZBBfH)
    
    // Test get() for type info
    val_uUEGnk = get("hello")
    pt(val_uUEGnk)
    

    m_oEPjqT = math.sqrt(56)
    h_TlfVag = hash.md5("test_OGYuWl")
    log.inf("hash:", h_TlfVag)
    pt(m_oEPjqT)
    

    asn fn async_task_KOkRmw(x) {
        rev x * 2
    }
    uni async_task_KOkRmw(25)
    
    asn fn worker_tjTcFF(y) {
        pt("working on", y)
    }
    mul 2 worker_tjTcFF(39)
    
    awt
    

    arr_hHSsoy = [52, 22+1, 48+2]
    map_tcUxVO = {"a": arr_hHSsoy[0], "b": arr_hHSsoy[1]}
    pt(map_tcUxVO["a"])
    

    use "dummy.y"
    
    e_YpDJQd = enc.b64("hello drylang")
    pt(e_YpDJQd)
    
    j_EEVcQj = json(`{"test": 123}`)
    pt(j_EEVcQj)
    
    // Test get() for type info
    val_gYVIHv = get("hello")
    pt(val_gYVIHv)
    

    cl Base_GPzRdi {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_jForGy <- Base_GPzRdi {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_WtVUfL = Child_jForGy(48, 40)
    pt(obj_WtVUfL.get_id())
    
