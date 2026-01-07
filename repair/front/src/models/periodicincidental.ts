import request from '~/global/request'

export default class Periodicincidental {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicincidental',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodicincidental',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicincidental',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicincidental/batch',
            data: item
        })

        return res
    }    

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodicincidental',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicincidental/${id}`
        })

        return res
    }

    static async getByPeriodic(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicincidental/get/periodic/${id}`
        })

        return res
    }    
}
