
    sum_WjexRq = 0
    lp 12 {
        sum_WjexRq = sum_WjexRq + 1
        if (sum_WjexRq > 100) { done }
    }
    pt(sum_WjexRq)
    

    try {
        throw_VWpQVF = unknown
        throw_VWpQVF()
    } err(e) {
        pt("caught error")
    }
    

    sum_MIhOdJ = 0
    lp 5 {
        sum_MIhOdJ = sum_MIhOdJ + 1
        if (sum_MIhOdJ > 100) { done }
    }
    pt(sum_MIhOdJ)
    

    arr_wMuIWD = [65, 45+1, 21+2]
    map_qbcUFc = {"a": arr_wMuIWD[0], "b": arr_wMuIWD[1]}
    pt(map_qbcUFc["a"])
    

    sum_PCfwPF = 0
    lp 14 {
        sum_PCfwPF = sum_PCfwPF + 1
        if (sum_PCfwPF > 100) { done }
    }
    pt(sum_PCfwPF)
    

    use "dummy.y"
    
    e_tqYgCv = enc.b64("hello drylang")
    pt(e_tqYgCv)
    
    j_isJiyn = json(`{"test": 123}`)
    pt(j_isJiyn)
    
    // Test get() for type info
    val_DKfwsE = get("hello")
    pt(val_DKfwsE)
    

    asn fn async_task_znULXn(x) {
        rev x * 2
    }
    uni async_task_znULXn(15)
    
    asn fn worker_fWuzKY(y) {
        pt("working on", y)
    }
    mul 2 worker_fWuzKY(65)
    
    awt
    

    m_MgcGWV = math.sqrt(48)
    h_MdxJpF = hash.md5("test_tPcrXp")
    log.inf("hash:", h_MdxJpF)
    pt(m_MgcGWV)
    

    sum_KXGIlM = 0
    lp 9 {
        sum_KXGIlM = sum_KXGIlM + 1
        if (sum_KXGIlM > 100) { done }
    }
    pt(sum_KXGIlM)
    

    cl Base_DnYvjh {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_QwhRoE <- Base_DnYvjh {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_dzRSiB = Child_QwhRoE(9, 23)
    pt(obj_dzRSiB.get_id())
    

    use "dummy.y"
    
    e_hfOzEQ = enc.b64("hello drylang")
    pt(e_hfOzEQ)
    
    j_JRUuZU = json(`{"test": 123}`)
    pt(j_JRUuZU)
    
    // Test get() for type info
    val_FdPGBH = get("hello")
    pt(val_FdPGBH)
    

    sum_VuuGGh = 0
    lp 6 {
        sum_VuuGGh = sum_VuuGGh + 1
        if (sum_VuuGGh > 100) { done }
    }
    pt(sum_VuuGGh)
    

    m_oEvnoG = math.sqrt(92)
    h_TqGFhI = hash.md5("test_pREOJG")
    log.inf("hash:", h_TqGFhI)
    pt(m_oEvnoG)
    

    arr_xLXzou = [82, 11+1, 49+2]
    map_sDoyyM = {"a": arr_xLXzou[0], "b": arr_xLXzou[1]}
    pt(map_sDoyyM["a"])
    
