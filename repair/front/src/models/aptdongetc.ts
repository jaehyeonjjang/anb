import request from '~/global/request'

export default class Aptdongetc {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/aptdongetc',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/aptdongetc',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/aptdongetc',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/aptdongetc',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/aptdongetc/${id}`
        })

        return res
    }

    static async countByAptdong(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/aptdongetc/count/aptdong/${id}`
        })

        return res
    }

    static async removeByAptdong(id: number) {
        let item = {
            id: id
        }
        
        const res = await request({
            method: 'DELETE',
            url: '/api/aptdongetc/byaptdong',
            data: item
        })

        return res
    }

    
}
