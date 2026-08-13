
    cl Base_rHbKcM {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_fJOAZD <- Base_rHbKcM {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_lGWniR = Child_fJOAZD(34, 42)
    pt(obj_lGWniR.get_id())
    

    fn calc_PzhPUV(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_Kxbahb = calc_PzhPUV(15, 9+1)
    pt(res_Kxbahb)
    

    fn calc_AoqDsq(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_PZzHPj = calc_AoqDsq(1, 1+1)
    pt(res_PZzHPj)
    

    cns C_xdMZTw = 15
    v_fNpSae = unknown
    v_fNpSae = C_xdMZTw + 86
    if (v_fNpSae > 0) {
        v_fNpSae = v_fNpSae * 2
    } el {
        v_fNpSae = 0
    }
    pt(v_fNpSae)
    

    cl Base_opcgHF {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_Zkexaj <- Base_opcgHF {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_efxedE = Child_Zkexaj(5, 92)
    pt(obj_efxedE.get_id())
    

    try {
        throw_xbpdRq = unknown
        throw_xbpdRq()
    } err(e) {
        pt("caught error")
    }
    
