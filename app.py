from flask import Flask, jsonify, request
from flask_cors import CORS
import os

app = Flask(__name__)
CORS(app)  # Enable CORS for all routes

@app.route('/')
def index():
    """Return basic API information."""
    return jsonify({
        'name': 'Nicholai\'s Personal API',
        'version': '1.0.0',
        'description': 'A personal API for Nicholai',
        'endpoints': [
            {'path': '/', 'method': 'GET', 'description': 'API information'},
            {'path': '/about', 'method': 'GET', 'description': 'About Nicholai'},
            {'path': '/projects', 'method': 'GET', 'description': 'List of projects'},
            {'path': '/contact', 'method': 'GET', 'description': 'Contact information'}
        ]
    })

@app.route('/about')
def about():
    """General information about me."""
    return jsonify({
        'name': 'Nicholai',
        'bio': 'Senior Remote Software Engineer',
        'skills': ['Skill 1', 'Skill 2', 'Skill 3'],
        'interests': ['Yoga', 'Livestreaming', 'Social Outreach']
    })


@app.route('/experience')
def experience():
    """My work experience."""
    return jsonify([
        {
            'title': 'Senior Remote Software Engineer',
            'company': 'Zoomph',
            'location': 'Remote, Virginia',
            'start_date': '6/2015',
            'active': True,
            'end_date': None,
            'tags': ['salaried'],
            'description': 'I am a senior remote software engineer at Zoomph, where I work with the frontend team to build and maintain their flagship product. I also work with integrating frontend features into the backend, so I work full-stack.'
        },
        {
            'title': 'Hot Yoga Instructor',
            'company': 'Corepower Yoga',
            'location': 'Northern Virginia',
            'start_date': 'Start Date 2',
            'active': True,
            'end_date': None,
            'tags': ['hourly'],
            'description': 'I am an internationally certified hot yoga instructor, and I teach twice-weekly classes in 100+ degree rooms.'
        },
        {
            'title': 'Social Outreach Coordinator',
            'company': 'CDCN',
            'location': 'Northern Virginia',
            'start_date': '2008',
            'active': False,
            'end_date': '2012',
            'tags': ['hourly'],
            'description': ''
        },
        {
            'title': 'Barista',
            'company': 'Borjo Coffee',
            'location': 'Norfolk, Virginia',
            'start_date': 'Start Date 4',
            'active': False,
            'end_date': '5/2015',
            'tags': ['hourly'],
            'description': 'I made coffee, made food, and helped organize concerts. The coffee shop was a popular spot for locals so I have fond memories of working there.'
        },
        {
            'title': 'Sales Associate',
            'company': 'Autozone',
            'location': 'Warrenton, Virginia',
            'start_date': '1/2011',
            'active': False,
            'end_date': '4/2011',
            'tags': ['hourly'],
            'description': ''
        },
        {
            'title': 'Sales Associate',
            'company': 'Play N\' Trade Videogames',
            'location': 'Warrenton, Virginia',
            'start_date': '6/2010',
            'active': False,
            'end_date': '11/2010',
            'tags': ['hourly'],
            'description': ''
        },
        {
            'title': 'Barista',
            'company': 'Starbucks Coffee',
            'location': 'Warrenton, Virginia',
            'start_date': '2007',
            'active': False,
            'end_date': '2008',
            'tags': ['hourly'],
            'description': ''
        },
        {
            'title': 'Associate',
            'company': 'Panera Bread',
            'location': 'Warrenton, Virginia',
            'start_date': '2006',
            'active': False,
            'end_date': '2008',
            'tags': ['hourly'],
            'description': ''
        },
    ])

@app.route('/projects')
def projects():
    """Projects i am working on."""
    return jsonify([
        {
            'name': 'Project 1',
            'description': 'Description of Project 1',
            'technologies': ['Tech 1', 'Tech 2'],
            'url': 'https://project1-url.com',
            'github': 'https://github.com/nicholai/project1'
        },
        {
            'name': 'Project 2',
            'description': 'Description of Project 2',
            'technologies': ['Tech 3', 'Tech 4'],
            'url': 'https://project2-url.com',
            'github': 'https://github.com/nicholai/project2'
        }
    ])

@app.route('/contact')
def contact():
    """Return contact information."""
    return jsonify({
        'email': 'nicholai@comfy.sh',
        'github': 'https://github.com/n3k0lai',
        'twitter': 'https://x.com/n3k0lai',
        'twitch': 'https://twitch.tv/n3k0lai',
        'spotify': 'https://open.spotify.com/user/n3k0lai',
        'linkedin': 'https://linkedin.com/in/nicholai',
        'instagram': 'https://instagram.com/n3k0lai',
    })

@app.errorhandler(404)
def not_found(e):
    """Handle 404 errors."""
    return jsonify({
        'error': 'Not found',
        'message': 'The requested URL was not found on this server.',
        'available_endpoints': [
            '/', '/about', '/projects', '/contact'
        ]
    }), 404

@app.errorhandler(500)
def server_error(e):
    """Handle 500 errors."""
    return jsonify({
        'error': 'Server error',
        'message': 'An internal server error occurred.'
    }), 500

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 5000))
    app.run(host='0.0.0.0', port=port, debug=False)
