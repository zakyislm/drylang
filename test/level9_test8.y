
    asn fn async_task_BfLfJS(x) {
        rev x * 2
    }
    uni async_task_BfLfJS(12)
    
    asn fn worker_nSelAU(y) {
        pt("working on", y)
    }
    mul 2 worker_nSelAU(37)
    
    awt
    

    sum_XVUhuf = 0
    lp 12 {
        sum_XVUhuf = sum_XVUhuf + 1
        if (sum_XVUhuf > 100) { done }
    }
    pt(sum_XVUhuf)
    

    asn fn async_task_WFzXaT(x) {
        rev x * 2
    }
    uni async_task_WFzXaT(3)
    
    asn fn worker_ZMLoGT(y) {
        pt("working on", y)
    }
    mul 2 worker_ZMLoGT(75)
    
    awt
    

    sum_BpchVv = 0
    lp 10 {
        sum_BpchVv = sum_BpchVv + 1
        if (sum_BpchVv > 100) { done }
    }
    pt(sum_BpchVv)
    

    arr_FQuDKC = [62, 61+1, 87+2]
    map_yzMvlD = {"a": arr_FQuDKC[0], "b": arr_FQuDKC[1]}
    pt(map_yzMvlD["a"])
    

    cns C_UvHEqD = 65
    v_viwssI = unknown
    v_viwssI = C_UvHEqD + 40
    if (v_viwssI > 0) {
        v_viwssI = v_viwssI * 2
    } el {
        v_viwssI = 0
    }
    pt(v_viwssI)
    

    cl Base_uvHwAK {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_RgBEkU <- Base_uvHwAK {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_DionnR = Child_RgBEkU(52, 71)
    pt(obj_DionnR.get_id())
    

    asn fn async_task_TYJaHz(x) {
        rev x * 2
    }
    uni async_task_TYJaHz(23)
    
    asn fn worker_ExVzfJ(y) {
        pt("working on", y)
    }
    mul 2 worker_ExVzfJ(64)
    
    awt
    

    asn fn async_task_sSjgKK(x) {
        rev x * 2
    }
    uni async_task_sSjgKK(96)
    
    asn fn worker_JSNHLs(y) {
        pt("working on", y)
    }
    mul 2 worker_JSNHLs(46)
    
    awt
    

    asn fn async_task_wnCpry(x) {
        rev x * 2
    }
    uni async_task_wnCpry(10)
    
    asn fn worker_HSAUCZ(y) {
        pt("working on", y)
    }
    mul 2 worker_HSAUCZ(100)
    
    awt
    

    asn fn async_task_gnwvaO(x) {
        rev x * 2
    }
    uni async_task_gnwvaO(59)
    
    asn fn worker_YlqkrC(y) {
        pt("working on", y)
    }
    mul 2 worker_YlqkrC(71)
    
    awt
    

    val_wuKdqD = 43 % 3
    on (val_wuKdqD) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_wuKdqD)
    

    asn fn async_task_pViQCV(x) {
        rev x * 2
    }
    uni async_task_pViQCV(7)
    
    asn fn worker_EKOtEy(y) {
        pt("working on", y)
    }
    mul 2 worker_EKOtEy(9)
    
    awt
    

    use "dummy.y"
    
    e_JHzdpt = enc.b64("hello drylang")
    pt(e_JHzdpt)
    
    j_xKWOVt = json(`{"test": 123}`)
    pt(j_xKWOVt)
    
    // Test get() for type info
    val_VuAElU = get("hello")
    pt(val_VuAElU)
    

    m_DqkGdJ = math.sqrt(71)
    h_cifAtl = hash.md5("test_GYjUmk")
    log.inf("hash:", h_cifAtl)
    pt(m_DqkGdJ)
    

    try {
        throw_evmGKh = unknown
        throw_evmGKh()
    } err(e) {
        pt("caught error")
    }
    
