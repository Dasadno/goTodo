   
async function fetchTasks() {
    try {
        const response = await fetch('/tasks');
        const tasks = await response.json();

        const list = document.getElementById('tasksList');
        list.innerHTML = '';

        tasks.forEach(task => {
            const li = document.createElement('li');
            li.classList.add('taskItem');

            const span = document.createElement('span');
            span.textContent = task.title;

            const btn = document.createElement('button');
            btn.className = 'deleteBtn';
            btn.innerHTML = '🗑️';
            btn.onclick = async () => {
        try {
            const response = await fetch(`/tasks/${task.id}`, {
                method: 'DELETE'
            });

            if (response.ok) {
                fetchTasks(); 
            } else {
                console.error('Ошибка при удалении:', response.status);
            }
        } catch (err) {
            console.error('Ошибка при удалении задачи:', err);
        }
        };
            li.appendChild(span);
            li.appendChild(btn);
            list.appendChild(li);
        });
    } catch (err) {
        console.error("Ошибка при загрузке задач:", err);
    }
}


const form = document.getElementById('taskForm');
  const input = document.getElementById('taskInput'); 

  form.addEventListener('submit', async (e) => {
    e.preventDefault();

    const title = input.value.trim(); 
    if (!title) return;

    try {
      const response = await fetch('/tasks', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          title: title,      
          content: "",
          checkMark: false
        })
      });

      if (response.ok) {
        input.value = "";     
        fetchTasks();         
      } else {
        console.error("Сервер вернул ошибку:", response.status);
      }
    } catch (err) {
      console.error("Ошибка при добавлении задачи:", err);
    }
  });

  fetchTasks();