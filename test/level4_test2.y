
    try {
        throw_BnnCwU = unknown
        throw_BnnCwU()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_TmctcZ {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_MGIeUi <- Base_TmctcZ {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_GLnjmc = Child_MGIeUi(34, 77)
    pt(obj_GLnjmc.get_id())
    

    fn calc_nhVJbR(x, y) {
        if (x <= y) { rev t }
        elif (x > y) { rev f }
        el { rev unknown }
    }
    res_tXZnnI = calc_nhVJbR(73, 76+1)
    pt(res_tXZnnI)
    

    val_brqvhx = 21 % 3
    on (val_brqvhx) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_brqvhx)
    

    val_qyrvRi = 82 % 3
    on (val_qyrvRi) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_qyrvRi)
    

    val_UWeUtz = 7 % 3
    on (val_UWeUtz) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_UWeUtz)
    
