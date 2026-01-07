import request from '~/global/request'

export default class Periodicpast {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicpast',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodicpast',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicpast',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodicpast',
            params: params
        })

        if (res.items == null) {
            res.items = []
        }

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicpast/${id}`
        })

        return res
    }
}
