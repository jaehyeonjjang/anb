import request from '~/global/request'

export default class Oldapt {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/oldapt',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/oldapt',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/oldapt',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/oldapt',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/oldapt/${id}`
        })

        return res
    }

    static async convert(items: any) {
        let data = {
            items: items
        }

        console.log(data)
        
        const res = await request({
            method: 'POST',
            url: '/api/oldapt/convert',
            data: data
        })

        return res
    }    
}
