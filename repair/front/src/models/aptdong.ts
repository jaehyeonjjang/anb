import request from '~/global/request'

export default class Aptdong {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/aptdong',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/aptdong',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/aptdong',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/aptdong',
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
            url: `/api/aptdong/${id}`
        })

        return res
    }

    static async blueprint(id: number, aptdong: number, items) {
        let item = {
            apt: id,
            aptdong: aptdong,
            items: items
        }
        const res = await request({
            method: 'POST',
            url: '/api/aptdong/blueprint',
            data: item
        })

        return res
    }  
}
