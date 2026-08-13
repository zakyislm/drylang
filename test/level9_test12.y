
    cl Base_vvYhLY {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_RDpBKO <- Base_vvYhLY {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_eAjkFA = Child_RDpBKO(78, 76)
    pt(obj_eAjkFA.get_id())
    

    val_VxGMak = 95 % 3
    on (val_VxGMak) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_VxGMak)
    

    asn fn async_task_IEzEic(x) {
        rev x * 2
    }
    uni async_task_IEzEic(88)
    
    asn fn worker_AsRytp(y) {
        pt("working on", y)
    }
    mul 2 worker_AsRytp(84)
    
    awt
    

    cns C_tkiGqd = 21
    v_yKuxgQ = unknown
    v_yKuxgQ = C_tkiGqd + 53
    if (v_yKuxgQ > 0) {
        v_yKuxgQ = v_yKuxgQ * 2
    } el {
        v_yKuxgQ = 0
    }
    pt(v_yKuxgQ)
    

    asn fn async_task_RtvuPb(x) {
        rev x * 2
    }
    uni async_task_RtvuPb(3)
    
    asn fn worker_ATBGYf(y) {
        pt("working on", y)
    }
    mul 2 worker_ATBGYf(21)
    
    awt
    

    arr_gIaOuo = [40, 18+1, 3+2]
    map_lIgIZU = {"a": arr_gIaOuo[0], "b": arr_gIaOuo[1]}
    pt(map_lIgIZU["a"])
    

    use "dummy.y"
    
    e_qWMXjc = enc.b64("hello drylang")
    pt(e_qWMXjc)
    
    j_SivdzP = json(`{"test": 123}`)
    pt(j_SivdzP)
    
    // Test get() for type info
    val_qjrdBU = get("hello")
    pt(val_qjrdBU)
    

    sum_NcjBBQ = 0
    lp 15 {
        sum_NcjBBQ = sum_NcjBBQ + 1
        if (sum_NcjBBQ > 100) { done }
    }
    pt(sum_NcjBBQ)
    

    cns C_hohQPK = 42
    v_WXwlCG = unknown
    v_WXwlCG = C_hohQPK + 29
    if (v_WXwlCG > 0) {
        v_WXwlCG = v_WXwlCG * 2
    } el {
        v_WXwlCG = 0
    }
    pt(v_WXwlCG)
    

    sum_RfbsvK = 0
    lp 10 {
        sum_RfbsvK = sum_RfbsvK + 1
        if (sum_RfbsvK > 100) { done }
    }
    pt(sum_RfbsvK)
    

    sum_XUAxMk = 0
    lp 8 {
        sum_XUAxMk = sum_XUAxMk + 1
        if (sum_XUAxMk > 100) { done }
    }
    pt(sum_XUAxMk)
    

    cl Base_fiXWWm {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_nKcaPl <- Base_fiXWWm {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_HmrpYZ = Child_nKcaPl(51, 88)
    pt(obj_HmrpYZ.get_id())
    

    asn fn async_task_zejaUi(x) {
        rev x * 2
    }
    uni async_task_zejaUi(25)
    
    asn fn worker_eWLhVR(y) {
        pt("working on", y)
    }
    mul 2 worker_eWLhVR(43)
    
    awt
    

    cns C_cJkjuQ = 37
    v_kZfyIf = unknown
    v_kZfyIf = C_cJkjuQ + 86
    if (v_kZfyIf > 0) {
        v_kZfyIf = v_kZfyIf * 2
    } el {
        v_kZfyIf = 0
    }
    pt(v_kZfyIf)
    

    try {
        throw_fIqadF = unknown
        throw_fIqadF()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_ILRJvH = unknown
        throw_ILRJvH()
    } err(e) {
        pt("caught error")
    }
    
