
    try {
        throw_vlvQcc = unknown
        throw_vlvQcc()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_wweWka {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_YhsYMT <- Base_wweWka {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_wMtAVu = Child_YhsYMT(2, 66)
    pt(obj_wMtAVu.get_id())
    

    cl Base_nhmeoX {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_MMEuuf <- Base_nhmeoX {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_CPsYFX = Child_MMEuuf(50, 84)
    pt(obj_CPsYFX.get_id())
    

    cns C_tgIBgT = 85
    v_ZJlPwx = unknown
    v_ZJlPwx = C_tgIBgT + 89
    if (v_ZJlPwx > 0) {
        v_ZJlPwx = v_ZJlPwx * 2
    } el {
        v_ZJlPwx = 0
    }
    pt(v_ZJlPwx)
    
