import request from '~/global/request'

export default class Breakdownhistory {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/breakdownhistory',
            data: item
        })

        return res
    }

    static async insertbatch(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/breakdownhistory/batch',
            data: item
        })

        return res
    }
    
    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/breakdownhistory',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/breakdownhistory',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/breakdownhistory/batch',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/breakdownhistory',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdownhistory/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdownhistory/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdownhistory/find/apt/${apt}`
        })

        return res
    }

    static async auto(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/breakdownhistory/auto/${apt}`
        })

        return res
    }    
}
