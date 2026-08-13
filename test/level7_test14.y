
    m_NNIvLp = math.sqrt(76)
    h_SnAlBr = hash.md5("test_IQjCAG")
    log.inf("hash:", h_SnAlBr)
    pt(m_NNIvLp)
    

    fn calc_FSCafT(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_otdyJS = calc_FSCafT(39, 80+1)
    pt(res_otdyJS)
    

    use "dummy.y"
    
    e_VhigKl = enc.b64("hello drylang")
    pt(e_VhigKl)
    
    j_orJsdp = json(`{"test": 123}`)
    pt(j_orJsdp)
    
    // Test get() for type info
    val_DeyQvi = get("hello")
    pt(val_DeyQvi)
    

    cl Base_qyPCnD {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_vneCXb <- Base_qyPCnD {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_CyGTOm = Child_vneCXb(30, 34)
    pt(obj_CyGTOm.get_id())
    

    cns C_hRpTsm = 79
    v_fInaJt = unknown
    v_fInaJt = C_hRpTsm + 90
    if (v_fInaJt > 0) {
        v_fInaJt = v_fInaJt * 2
    } el {
        v_fInaJt = 0
    }
    pt(v_fInaJt)
    

    sum_HIISBY = 0
    lp 5 {
        sum_HIISBY = sum_HIISBY + 1
        if (sum_HIISBY > 100) { done }
    }
    pt(sum_HIISBY)
    

    cl Base_YCFrMv {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_wrqoHQ <- Base_YCFrMv {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_oBFLRG = Child_wrqoHQ(23, 82)
    pt(obj_oBFLRG.get_id())
    

    arr_rrLlyZ = [99, 68+1, 32+2]
    map_FRnigT = {"a": arr_rrLlyZ[0], "b": arr_rrLlyZ[1]}
    pt(map_FRnigT["a"])
    

    try {
        throw_wpYgcQ = unknown
        throw_wpYgcQ()
    } err(e) {
        pt("caught error")
    }
    

    cns C_EERDvo = 74
    v_psudwz = unknown
    v_psudwz = C_EERDvo + 50
    if (v_psudwz > 0) {
        v_psudwz = v_psudwz * 2
    } el {
        v_psudwz = 0
    }
    pt(v_psudwz)
    

    sum_nDWnIZ = 0
    lp 13 {
        sum_nDWnIZ = sum_nDWnIZ + 1
        if (sum_nDWnIZ > 100) { done }
    }
    pt(sum_nDWnIZ)
    

    asn fn async_task_YSWBfY(x) {
        rev x * 2
    }
    uni async_task_YSWBfY(40)
    
    asn fn worker_wOkscO(y) {
        pt("working on", y)
    }
    mul 2 worker_wOkscO(27)
    
    awt
    

    asn fn async_task_QDXAOw(x) {
        rev x * 2
    }
    uni async_task_QDXAOw(76)
    
    asn fn worker_qiHcZm(y) {
        pt("working on", y)
    }
    mul 2 worker_qiHcZm(47)
    
    awt
    

    asn fn async_task_BpmDbv(x) {
        rev x * 2
    }
    uni async_task_BpmDbv(62)
    
    asn fn worker_ERvzHB(y) {
        pt("working on", y)
    }
    mul 2 worker_ERvzHB(66)
    
    awt
    
