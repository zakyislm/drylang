
    sum_LcQhGW = 0
    lp 6 {
        sum_LcQhGW = sum_LcQhGW + 1
        if (sum_LcQhGW > 100) { done }
    }
    pt(sum_LcQhGW)
    

    cns C_ILCnLG = 76
    v_hvfewx = unknown
    v_hvfewx = C_ILCnLG + 77
    if (v_hvfewx > 0) {
        v_hvfewx = v_hvfewx * 2
    } el {
        v_hvfewx = 0
    }
    pt(v_hvfewx)
    

    val_YKzsGv = 75 % 3
    on (val_YKzsGv) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_YKzsGv)
    

    cns C_KmygbQ = 60
    v_CJfnbK = unknown
    v_CJfnbK = C_KmygbQ + 52
    if (v_CJfnbK > 0) {
        v_CJfnbK = v_CJfnbK * 2
    } el {
        v_CJfnbK = 0
    }
    pt(v_CJfnbK)
    

    try {
        throw_GzFNxw = unknown
        throw_GzFNxw()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_TTNcqn {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_AjGmkR <- Base_TTNcqn {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_unNEfN = Child_AjGmkR(79, 19)
    pt(obj_unNEfN.get_id())
    

    cl Base_iAtlmT {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_oPmseC <- Base_iAtlmT {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_SjufEw = Child_oPmseC(7, 63)
    pt(obj_SjufEw.get_id())
    

    m_kdPJbF = math.sqrt(25)
    h_kbnxgc = hash.md5("test_FCIyqc")
    log.inf("hash:", h_kbnxgc)
    pt(m_kdPJbF)
    

    arr_HinlHe = [96, 13+1, 65+2]
    map_ejaNbJ = {"a": arr_HinlHe[0], "b": arr_HinlHe[1]}
    pt(map_ejaNbJ["a"])
    

    sum_kVnnLI = 0
    lp 5 {
        sum_kVnnLI = sum_kVnnLI + 1
        if (sum_kVnnLI > 100) { done }
    }
    pt(sum_kVnnLI)
    

    val_CBhzyG = 57 % 3
    on (val_CBhzyG) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_CBhzyG)
    

    cl Base_AMJTIR {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_WSTRHz <- Base_AMJTIR {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_moGiYU = Child_WSTRHz(11, 64)
    pt(obj_moGiYU.get_id())
    

    m_AYaGSe = math.sqrt(23)
    h_Dganlh = hash.md5("test_xzoVuk")
    log.inf("hash:", h_Dganlh)
    pt(m_AYaGSe)
    

    val_JOzTJX = 45 % 3
    on (val_JOzTJX) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_JOzTJX)
    

    cl Base_UJsIqb {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_mjSieB <- Base_UJsIqb {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_CutQnT = Child_mjSieB(94, 88)
    pt(obj_CutQnT.get_id())
    

    asn fn async_task_TlSrPW(x) {
        rev x * 2
    }
    uni async_task_TlSrPW(34)
    
    asn fn worker_taGAJc(y) {
        pt("working on", y)
    }
    mul 2 worker_taGAJc(33)
    
    awt
    
