
    cl Base_CTbXgS {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_FKyNse <- Base_CTbXgS {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_dAyApE = Child_FKyNse(44, 1)
    pt(obj_dAyApE.get_id())
    

    cns C_LbBFEK = 17
    v_IXEEoT = unknown
    v_IXEEoT = C_LbBFEK + 61
    if (v_IXEEoT > 0) {
        v_IXEEoT = v_IXEEoT * 2
    } el {
        v_IXEEoT = 0
    }
    pt(v_IXEEoT)
    

    fn calc_ouXcrP(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_ghlmtT = calc_ouXcrP(49, 24+1)
    pt(res_ghlmtT)
    

    sum_DBmgMb = 0
    lp 14 {
        sum_DBmgMb = sum_DBmgMb + 1
        if (sum_DBmgMb > 100) { done }
    }
    pt(sum_DBmgMb)
    
