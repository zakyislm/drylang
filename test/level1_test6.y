
    try {
        throw_qnmmsi = unknown
        throw_qnmmsi()
    } err(e) {
        pt("caught error")
    }
    

    cl Base_vToXDX {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_dKzJuN <- Base_vToXDX {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_ItDOkk = Child_dKzJuN(88, 88)
    pt(obj_ItDOkk.get_id())
    

    sum_EweQTy = 0
    lp 6 {
        sum_EweQTy = sum_EweQTy + 1
        if (sum_EweQTy > 100) { done }
    }
    pt(sum_EweQTy)
    
