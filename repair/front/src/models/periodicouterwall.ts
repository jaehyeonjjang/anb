import request from '~/global/request'

export default class Periodicouterwall {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicouterwall',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodicouterwall',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicouterwall',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicouterwall/batch',
            data: item
        })

        return res
    }    

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodicouterwall',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicouterwall/${id}`
        })

        return res
    }

    static async getByPeriodic(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicouterwall/get/periodic/${id}`
        })

        return res
    }
}
