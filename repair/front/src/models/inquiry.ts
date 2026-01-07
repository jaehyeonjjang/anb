import request from '~/global/request'

export default class Inquiry {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/inquiry',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/inquiry',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/inquiry',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/inquiry/batch',
            data: item
        })

        return res
    }    

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/inquiry',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/inquiry/${id}`
        })

        return res
    }
}
