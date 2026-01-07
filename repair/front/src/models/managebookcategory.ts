import request from '~/global/request'

export default class Managebookcategory {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/managebookcategory',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/managebookcategory',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/managebookcategory',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/managebookcategory/batch',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/managebookcategory',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/managebookcategory/${id}`
        })

        return res
    }

    static async process(id: number, filename: string) {
        const item = {
            id: id,
            filename: filename
        }
        
        const res = await request({
            method: 'POST',
            url: `/api/managebookcategory/process`,
            data: item
        })

        return res
    }
}
