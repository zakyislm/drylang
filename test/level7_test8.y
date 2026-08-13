
    try {
        throw_pjaWMO = unknown
        throw_pjaWMO()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_kpQCrG(x) {
        rev x * 2
    }
    uni async_task_kpQCrG(29)
    
    asn fn worker_MHRNTf(y) {
        pt("working on", y)
    }
    mul 2 worker_MHRNTf(13)
    
    awt
    

    try {
        throw_jNUXjd = unknown
        throw_jNUXjd()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_jjbdEZ = unknown
        throw_jjbdEZ()
    } err(e) {
        pt("caught error")
    }
    

    asn fn async_task_GdnhHQ(x) {
        rev x * 2
    }
    uni async_task_GdnhHQ(78)
    
    asn fn worker_sxJReY(y) {
        pt("working on", y)
    }
    mul 2 worker_sxJReY(82)
    
    awt
    

    val_nMfIUG = 93 % 3
    on (val_nMfIUG) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_nMfIUG)
    

    val_MIuFDg = 15 % 3
    on (val_MIuFDg) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_MIuFDg)
    

    cns C_UykgQL = 45
    v_lLTzPP = unknown
    v_lLTzPP = C_UykgQL + 72
    if (v_lLTzPP > 0) {
        v_lLTzPP = v_lLTzPP * 2
    } el {
        v_lLTzPP = 0
    }
    pt(v_lLTzPP)
    

    sum_JTjjxC = 0
    lp 10 {
        sum_JTjjxC = sum_JTjjxC + 1
        if (sum_JTjjxC > 100) { done }
    }
    pt(sum_JTjjxC)
    

    sum_BSAbzO = 0
    lp 9 {
        sum_BSAbzO = sum_BSAbzO + 1
        if (sum_BSAbzO > 100) { done }
    }
    pt(sum_BSAbzO)
    

    cl Base_xeNrgH {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_gSlBmQ <- Base_xeNrgH {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_xadyVw = Child_gSlBmQ(86, 51)
    pt(obj_xadyVw.get_id())
    

    try {
        throw_rVkiCC = unknown
        throw_rVkiCC()
    } err(e) {
        pt("caught error")
    }
    

    arr_SpMcbT = [40, 27+1, 53+2]
    map_ykbJYt = {"a": arr_SpMcbT[0], "b": arr_SpMcbT[1]}
    pt(map_ykbJYt["a"])
    

    val_HKpofM = 98 % 3
    on (val_HKpofM) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_HKpofM)
    
