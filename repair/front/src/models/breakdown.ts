import request from '~/global/request'

export default class Breakdown {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/breakdown',
            data: item
        })

        return res
    }

    static async insertbatch(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/breakdown/batch',
            data: item
        })

        return res
    }
    
    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/breakdown',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/breakdown',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/breakdown/batch',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/breakdown',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdown/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdown/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdown/find/apt/${apt}`
        })

        return res
    }

    static async countByAptDong(apt: number, dong: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdown/count/aptdong/${apt}?dong=${dong}`
        })

        return res.count
    }

    static async deduplication(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/breakdown/deduplication',
            data: item
        })

        return res
    }

    static async updateDuedateById(duedate: number, id: number) {
        let item = {
            id,
            duedate 
        }
        
        const res = await request({
            method: 'PUT',
            url: '/api/breakdown/duedatebyid',
            data: item
        })

        return res
    }

    static async updateLastdateById(lastdate: number, id: number) {
        let item = {
            id,
            lastdate 
        }
        
        const res = await request({
            method: 'PUT',
            url: '/api/breakdown/lastdatebyid',
            data: item
        })

        return res
    }    

    static async updateLastdate(date: number, ids: string) {
        let item = {            
            ids,
            date 
        }

        const res = await request({
            method: 'POST',
            url: '/api/breakdown/lastdate',
            data: item
        })

        return res
    }
    
    static async updateDuedate(apt: number, date: number, ids: string) {
        let item = {
            apt,
            ids,
            date 
        }

        const res = await request({
            method: 'POST',
            url: '/api/breakdown/duedate',
            data: item
        })

        return res
    }
}
