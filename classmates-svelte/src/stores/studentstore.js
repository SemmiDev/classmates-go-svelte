import { writable } from 'svelte/store';

export const students = writable([]);

const fetchstudents = async () => {
	const url = 'http://localhost:8080/api/v1/classmates/';
	const res = await fetch(url);
	const result = await res.json();
	const studentData = result.data.map(item => {
		return {
			id: item.id,
			name: item.name,
            identifier: item.identifier,
			email: item.email,
            image: item.image_url
		}
	})
	students.set(studentData);
};

fetchstudents();
