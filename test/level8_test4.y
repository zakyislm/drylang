
    cl Base_DfPvNQ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_IVATwM <- Base_DfPvNQ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_nqbazG = Child_IVATwM(66, 35)
    pt(obj_nqbazG.get_id())
    

    sum_MVKHAM = 0
    lp 13 {
        sum_MVKHAM = sum_MVKHAM + 1
        if (sum_MVKHAM > 100) { done }
    }
    pt(sum_MVKHAM)
    

    sum_xLtdxd = 0
    lp 12 {
        sum_xLtdxd = sum_xLtdxd + 1
        if (sum_xLtdxd > 100) { done }
    }
    pt(sum_xLtdxd)
    

    use "dummy.y"
    
    e_EpmUHv = enc.b64("hello drylang")
    pt(e_EpmUHv)
    
    j_tLZaJN = json(`{"test": 123}`)
    pt(j_tLZaJN)
    
    // Test get() for type info
    val_xuTYtV = get("hello")
    pt(val_xuTYtV)
    

    arr_fLRABb = [56, 10+1, 88+2]
    map_FuCkYt = {"a": arr_fLRABb[0], "b": arr_fLRABb[1]}
    pt(map_FuCkYt["a"])
    

    arr_ApAkWa = [48, 51+1, 8+2]
    map_uKDSYP = {"a": arr_ApAkWa[0], "b": arr_ApAkWa[1]}
    pt(map_uKDSYP["a"])
    

    arr_sPpvBf = [75, 89+1, 72+2]
    map_lIIpoW = {"a": arr_sPpvBf[0], "b": arr_sPpvBf[1]}
    pt(map_lIIpoW["a"])
    

    use "dummy.y"
    
    e_MNHdaj = enc.b64("hello drylang")
    pt(e_MNHdaj)
    
    j_nGrXBT = json(`{"test": 123}`)
    pt(j_nGrXBT)
    
    // Test get() for type info
    val_rXSYCj = get("hello")
    pt(val_rXSYCj)
    

    cl Base_hhjQqT {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_YCjQna <- Base_hhjQqT {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_DLwwPH = Child_YCjQna(76, 46)
    pt(obj_DLwwPH.get_id())
    

    cns C_FhNHdQ = 65
    v_oWKhcv = unknown
    v_oWKhcv = C_FhNHdQ + 5
    if (v_oWKhcv > 0) {
        v_oWKhcv = v_oWKhcv * 2
    } el {
        v_oWKhcv = 0
    }
    pt(v_oWKhcv)
    

    cl Base_lyDcpo {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_pwmSMB <- Base_lyDcpo {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_CUnzGW = Child_pwmSMB(17, 19)
    pt(obj_CUnzGW.get_id())
    

    arr_RlnnPn = [51, 91+1, 25+2]
    map_aInWMc = {"a": arr_RlnnPn[0], "b": arr_RlnnPn[1]}
    pt(map_aInWMc["a"])
    

    try {
        throw_mzMbpi = unknown
        throw_mzMbpi()
    } err(e) {
        pt("caught error")
    }
    

    fn calc_TqxFUe(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_kgAtmB = calc_TqxFUe(42, 94+1)
    pt(res_kgAtmB)
    

    fn calc_xnTTyG(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_DRCYOC = calc_xnTTyG(41, 21+1)
    pt(res_DRCYOC)
    
