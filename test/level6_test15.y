
    cns C_chwQyh = 96
    v_fRVPHE = unknown
    v_fRVPHE = C_chwQyh + 14
    if (v_fRVPHE > 0) {
        v_fRVPHE = v_fRVPHE * 2
    } el {
        v_fRVPHE = 0
    }
    pt(v_fRVPHE)
    

    arr_VWUCmC = [42, 55+1, 2+2]
    map_Ybfrib = {"a": arr_VWUCmC[0], "b": arr_VWUCmC[1]}
    pt(map_Ybfrib["a"])
    

    cns C_RCtHyv = 22
    v_pHNOUC = unknown
    v_pHNOUC = C_RCtHyv + 39
    if (v_pHNOUC > 0) {
        v_pHNOUC = v_pHNOUC * 2
    } el {
        v_pHNOUC = 0
    }
    pt(v_pHNOUC)
    

    fn calc_nWNHsf(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_OlYumh = calc_nWNHsf(69, 96+1)
    pt(res_OlYumh)
    

    try {
        throw_GDGcwR = unknown
        throw_GDGcwR()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_nQbNPJ = unknown
        throw_nQbNPJ()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_JvLYgs {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_NoViLN <- Base_JvLYgs {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_TUgjeW = Child_NoViLN(44, 3)
    pt(obj_TUgjeW.get_id())
    

    cl Base_swbuJo {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_vjCcvV <- Base_swbuJo {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_YTXHin = Child_vjCcvV(60, 88)
    pt(obj_YTXHin.get_id())
    
