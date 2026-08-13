
    val_yyxcrn = 22 % 3
    on (val_yyxcrn) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_yyxcrn)
    

    sum_FVmXSA = 0
    lp 10 {
        sum_FVmXSA = sum_FVmXSA + 1
        if (sum_FVmXSA > 100) { done }
    }
    pt(sum_FVmXSA)
    

    sum_JGVKrY = 0
    lp 6 {
        sum_JGVKrY = sum_JGVKrY + 1
        if (sum_JGVKrY > 100) { done }
    }
    pt(sum_JGVKrY)
    

    val_wIJuwm = 33 % 3
    on (val_wIJuwm) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_wIJuwm)
    

    m_wbOZDY = math.sqrt(75)
    h_ViBQXo = hash.md5("test_fxHaxo")
    log.inf("hash:", h_ViBQXo)
    pt(m_wbOZDY)
    

    m_gRoeVC = math.sqrt(14)
    h_mfgHae = hash.md5("test_UgKcMC")
    log.inf("hash:", h_mfgHae)
    pt(m_gRoeVC)
    

    cns C_PEHrjY = 6
    v_ekqZbF = unknown
    v_ekqZbF = C_PEHrjY + 97
    if (v_ekqZbF > 0) {
        v_ekqZbF = v_ekqZbF * 2
    } el {
        v_ekqZbF = 0
    }
    pt(v_ekqZbF)
    

    arr_oVGiie = [43, 38+1, 7+2]
    map_idGSAc = {"a": arr_oVGiie[0], "b": arr_oVGiie[1]}
    pt(map_idGSAc["a"])
    

    arr_aEoqsI = [48, 64+1, 74+2]
    map_LnvMoV = {"a": arr_aEoqsI[0], "b": arr_aEoqsI[1]}
    pt(map_LnvMoV["a"])
    

    cl Base_PhUxUT {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_qwzqBj <- Base_PhUxUT {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_rJFfwx = Child_qwzqBj(25, 30)
    pt(obj_rJFfwx.get_id())
    

    cns C_NYEsxH = 50
    v_KXYScy = unknown
    v_KXYScy = C_NYEsxH + 72
    if (v_KXYScy > 0) {
        v_KXYScy = v_KXYScy * 2
    } el {
        v_KXYScy = 0
    }
    pt(v_KXYScy)
    

    val_wBTKmk = 54 % 3
    on (val_wBTKmk) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_wBTKmk)
    

    try {
        throw_arjufP = unknown
        throw_arjufP()
    } err(e) {
        pt("caught error")
    }
    

    m_OebmUn = math.sqrt(75)
    h_SNSMsg = hash.md5("test_FMzBKr")
    log.inf("hash:", h_SNSMsg)
    pt(m_OebmUn)
    

    try {
        throw_ggXFmC = unknown
        throw_ggXFmC()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_FYweZn = unknown
        throw_FYweZn()
    } err(e) {
        pt("caught error")
    }
    
