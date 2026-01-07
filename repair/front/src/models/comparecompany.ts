import request from '~/global/request'

export default class Comparecompany {
    static defaults = [' ', '기본', '기본 아님']

    static getDefault(value: number) {
        return this.defaults[value]
    }

    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/comparecompany',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/comparecompany',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/comparecompany',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/comparecompany',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/comparecompany/${id}`
        })

        return res
    }
}
