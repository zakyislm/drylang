
    fn calc_wmcdXj(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_BXQdvo = calc_wmcdXj(56, 75+1)
    pt(res_BXQdvo)
    

    arr_mQrEPe = [12, 59+1, 94+2]
    map_xbIkYB = {"a": arr_mQrEPe[0], "b": arr_mQrEPe[1]}
    pt(map_xbIkYB["a"])
    

    cl Base_IRKquv {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_KzCuFd <- Base_IRKquv {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_MaWyaQ = Child_KzCuFd(44, 1)
    pt(obj_MaWyaQ.get_id())
    

    sum_zLhgAN = 0
    lp 10 {
        sum_zLhgAN = sum_zLhgAN + 1
        if (sum_zLhgAN > 100) { done }
    }
    pt(sum_zLhgAN)
    

    cl Base_PtZcUN {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_YJqWtV <- Base_PtZcUN {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_JluJgn = Child_YJqWtV(90, 69)
    pt(obj_JluJgn.get_id())
    

    try {
        throw_hGaase = unknown
        throw_hGaase()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_GaCqXS = unknown
        throw_GaCqXS()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_jRVoCg = unknown
        throw_jRVoCg()
    } err(e) {
        pt("caught error")
    }
    

    try {
        throw_mcNjyx = unknown
        throw_mcNjyx()
    } err(e) {
        pt("caught error")
    }
    

    use "dummy.y"
    
    e_ctzeIG = enc.b64("hello drylang")
    pt(e_ctzeIG)
    
    j_jNKrDE = json(`{"test": 123}`)
    pt(j_jNKrDE)
    
    // Test get() for type info
    val_OVGAwi = get("hello")
    pt(val_OVGAwi)
    

    cl Base_hKdzqG {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_mVmhys <- Base_hKdzqG {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_vtkQjD = Child_mVmhys(54, 48)
    pt(obj_vtkQjD.get_id())
    

    fn calc_QbSoWs(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_kzBtsw = calc_QbSoWs(64, 52+1)
    pt(res_kzBtsw)
    

    arr_QEMuel = [69, 36+1, 72+2]
    map_HObBqA = {"a": arr_QEMuel[0], "b": arr_QEMuel[1]}
    pt(map_HObBqA["a"])
    

    sum_WhcaAg = 0
    lp 6 {
        sum_WhcaAg = sum_WhcaAg + 1
        if (sum_WhcaAg > 100) { done }
    }
    pt(sum_WhcaAg)
    
