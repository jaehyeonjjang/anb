import request from '~/global/request'

export default class Reviewcontentbasic {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/reviewcontentbasic',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/reviewcontentbasic',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/reviewcontentbasic',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/reviewcontentbasic',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/reviewcontentbasic/${id}`
        })

        return res
    }
}
