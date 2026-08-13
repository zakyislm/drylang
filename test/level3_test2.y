
    sum_lutbPl = 0
    lp 10 {
        sum_lutbPl = sum_lutbPl + 1
        if (sum_lutbPl > 100) { done }
    }
    pt(sum_lutbPl)
    

    cl Base_BgMyny {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_kalnxF <- Base_BgMyny {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_GxYGYY = Child_kalnxF(33, 84)
    pt(obj_GxYGYY.get_id())
    

    val_gywHct = 36 % 3
    on (val_gywHct) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_gywHct)
    

    try {
        throw_PWsRhm = unknown
        throw_PWsRhm()
    } err(e) {
        pt("caught error")
    }
    

    cns C_PGbwIH = 14
    v_fQBKLN = unknown
    v_fQBKLN = C_PGbwIH + 61
    if (v_fQBKLN > 0) {
        v_fQBKLN = v_fQBKLN * 2
    } el {
        v_fQBKLN = 0
    }
    pt(v_fQBKLN)
    
